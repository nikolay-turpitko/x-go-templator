cwd=$(pwd)
cd /tmp
wget http://download.icu-project.org/files/icu4c/60.1/icu4c-60_1-src.tgz
tar -xf icu4c-60_1-src.tgz
cd icu/source
./configure --prefix=/usr && make
sudo make install
sudo chmod a+rx /usr/lib/icu/60.1
sudo chmod a+rx /usr/include/unicode
cd $cwd
