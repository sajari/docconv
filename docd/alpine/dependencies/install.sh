#!/usr/bin/env sh

mkdir /tidy
mkdir /unrtf

# Install Tidy
tar -C /tidy -xvf tidy-html5-5.4.0.tar.gz
cd /tidy/tidy-html5-5.4.0/build/cmake
cmake ../.. -DCMAKE_BUILD_TYPE=Release
make 
make install

# Install UnRTF
tar -C /unrtf -xvf /unrtf-0.21.9.tar.gz
cd /unrtf/unrtf-0.21.9
ls
./configure --help
./configure; make; make install; make clean

cd /
rm -r /tidy
rm -r /unrtf
rm *.tar.gz