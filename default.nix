{
  buildGoModule,
  lib,
  ...
}: let
  pname = "echo";
  version = "0.1.1";
in
  buildGoModule {
    inherit pname version;
    src = builtins.filterSource (path: type: type != "directory" || baseNameOf path != ".git" || lib.hasSuffix ".nix" path) ./.;
    vendorHash = null;

    ldflags = ["-s" "-w" "-X main.version=${version}"];

    meta = {
      description = "Simple & lighweight mock server on localhost for testing.";
      homepage = "https://github.com/notAShelf/echo";
      license = lib.licenses.gpl3Only;
      maintainers = with lib.maintainers; [NotAShelf];
      mainProgram = pname;
    };
  }
