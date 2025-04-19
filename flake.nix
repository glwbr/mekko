{
  description = "🍹 Juice - Brazilian NFC-e Invoice Extraction Library";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    { nixpkgs, ... }:
    let
      systems = [
        "aarch64-linux"
        "x86_64-linux"
      ];

      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f pkgsFor.${system});
      pkgsFor = nixpkgs.lib.genAttrs systems (
        system:
        import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        }
      );
    in
    {
      devShells = forEachSystem (pkgs: {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            golangci-lint
            treefmt
            nixfmt-rfc-style
          ];

          shellHook = ''
            echo "A juicy environment is ready!"
          '';
        };
      });
    };
}
