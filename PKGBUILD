pkgname="webpalm"
pkgver="v2.0.13"
pkgrel=1
pkgdesc="webpalm is a powerful command-line tool for website mapping and web scraping, making it an ideal tool for web scraping and data extraction."
arch=("x86_64" "i686" "aarch64")
url="https://malwarize.live/"
license="GPL3"
depends=""
makedepends="git"
source_x86_64="$pkgname-${pkgver}_linux_amd64.tar.gz::https://github.com/Malwarize/$pkgname/releases/download/$pkgver/webpalm_2.0.13_linux_amd64.tar.gz"
provides="webpalm"

package() {
	install -Dm755 "webpalm" "$pkgdir/usr/bin/webpalm"
}
