pkgname=waypower
pkgver=0.0.1
pkgrel=1
pkgdesc='Aggregrates Power Consumption for use in Waybar'
arch=('x86_64')
url="https://github.com/Garionion/$pkgname"
license=('GPL')
makedepends=('go')
source=("main.go")
sha256sums=('fd0d76d5e7d03900f9111c00a1b50f124d8ac2ee09666752bdb505ed3639c10e')

build() {
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
  go build -o $pkgname .
}

package() {
  install -Dm755 $pkgname "$pkgdir"/usr/bin/$pkgname
}
