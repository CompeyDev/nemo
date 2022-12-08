echo "[*] Installing gccgo toolchain..."

git clone --branch devel/gccgo git://gcc.gnu.org/git/gcc.git gccgo
mkdir data
cd data
../gccgo/configure --prefix=/opt/gccgo --enable-languages=c,c++,go --disable-multilib
make > /dev/null 2>&1
make install > /dev/null 2>&1