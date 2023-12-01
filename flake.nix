{
  description = "Simple & lighweight mock server on localhost for testing. ";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = ["x86_64-linux" "aarch64-linux"];
    forEachSystem = nixpkgs.lib.genAttrs systems;

    pkgsForEach = nixpkgs.legacyPackages;
  in {
    packages = forEachSystem (system: {
      echo = pkgsForEach.${system}.callPackage ./default.nix {};
      default = self.packages.${system}.echo;
    });

    devShells = forEachSystem (system: {
      default = pkgsForEach.${system}.callPackage ./shell.nix {};
    });
  };
}
