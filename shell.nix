{ pkgs ? import <nixpkgs> {} }:

let
  tunScript = pkgs.writeShellScript "run-tun2proxy" ''
    sudo nohup tun2proxy-bin --setup  --proxy socks5://127.0.0.1:1080 --bypass 0.0.0.0 --tun tun0  --dns virtual > /tmp/tun2proxy.log 2>&1 &
  '';
in
pkgs.mkShell {
  buildInputs = [
    pkgs.tun2proxy
    pkgs.xray
    pkgs.go
  ];

  shellHook = ''
    systemd-run --unit=xray-temp.service --description="Temporary Xray Service" --working-directory="$(pwd)" xray
    ${tunScript}
    cleanup() {
      echo "Stopping temporary services..."
      systemctl stop xray-temp.service
      sudo pkill -f "tun2proxy" 
    }
    trap cleanup EXIT
  '';
}
