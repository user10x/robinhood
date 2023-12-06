class Rhctl < Formula
  desc "Description of your software"
  homepage "https://github.com/robhinhood"
  version "v1.0.0"  # Specify the version here

  if OS.mac?
    url "https://github.com/user10x/robinhood/releases/download/v1.0.0/rhctl_darwin_arm64"
    sha256 "cfbe7b9caeb2e1c5d6f0685dd7f2449fa34f901dda8b5a25acdc4cc2ddc59ff5"
  elsif OS.linux?
    url "https://github.com/user10x/robinhood/releases/download/v1.0.0/rhctl_linux_amd64.tar.gz"
    sha256 "????"
  end

  def install
    bin.install "/usr/local/rhctl"
  end

  test do
    # Add test instructions if applicable
    system "#{bin}/rhctl", "--help"
  end
end
