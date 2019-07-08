require_relative '../lib/github_downloader'

class GoStarter < Formula
  desc "Go Project Starter"
  homepage "https://github.com/magento-mcom/go-starter"
  url "${URL}", :using => GithubAssetDownloadStrategy
  sha256 "${HASH}"
  version "${VERSION}"
  revision ${REVISION}

  def install
    bin.install "go-starter"
    bin.install "go-starter-replace"
    bin.install "go-starter-github"
    bin.install "go-starter-drone"
  end
end
