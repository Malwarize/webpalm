# Maintainer: XORbit <XORbit01@proton.me> Mercury <elhirchek@gmail.com>
maintainer=XORbit
pkgname_orig=webpalm
pkgname=webpalm-bin
pkgver="2.0.13"
pkgrel=1
pkgdesc="powerful command-line tool for website mapping and web scraping, making it an ideal tool for web scraping and data extraction"
arch=('x86_64')
url="https://github.com/Malwarize/webpalm"
provides=('webpam')
license=('GPL3')
conflicts=('webpalm' 'webpalm-git')
binname=${pkgname_orig}-${pkgver}-${pkgrel}
dirname=${pkgname_orig}_${pkgver}_linux_amd64
source_x86_64=(
	"${binname}.tar.gz::${url}/releases/download/v${pkgver}/${dirname}.tar.gz"
)
sha256sums_x86_64=(
	'a60005bd4008f377413f84fecb90bc8a0459f88e76b03df085796f6f3454cc0c'
)

package(){
	install -Dm755 "${pkgname_orig}" "$pkgdir/usr/bin/${pkgname_orig}"
}
