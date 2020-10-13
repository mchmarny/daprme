class Daprme < Formula
  desc "New app wizard for creation of a new Dapr project"
  homepage "https://dapr.io"
  url "https://github.com/dapr-templates/daprme/releases/download/v0.5.1/daprme"
  sha256 "20285d408f3d4a96c2004e4562923a8bff9375124fe047cd70f1a8e618006adf"
  license "MIT"

  def install
    bin.install "daprme" => "daprme"
  end

  test do
    system "#{bin}/daprme", "--version"
  end
end
