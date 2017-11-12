cwd=$(pwd)
cd /tmp
wget http://download.icu-project.org/files/icu4c/60.1/icu4c-60_1-src.tgz
tar -xf icu4c-60_1-src.tgz
cd icu/source
./configure --prefix=/usr && make
sudo make install
cd $cwd
