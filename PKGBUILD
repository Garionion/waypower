pkgname=waypower
pkgver=0.0.1
pkgrel=1
pkgdesc='Aggregrates Power Consumption for use in Waybar'
arch=('x86_64')
url="https://github.com/Garionion/$pkgname"
license=('GPL')
makedepends=('go')
source=("main.go")
sha256sums=('47cf4d6252a5ff3c1dfbb15ef838dee4d5c85daf9c84fa6918a9e650b8350a9f')

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
