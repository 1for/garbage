#!/bin/sh

#打印当前执行的命令
set -x

DOWNLOAD_DIR=~/download
INSTALL_DIR=~/program

EMACS_DOWNLOAD_URL="http://mirrors.ustc.edu.cn/gnu/emacs/emacs-26.3.tar.gz"
EMACS_DIR=emacs-26.3
EMACS_GZ=emacs-26.3.tar.gz

if [ ! -e $DOWNLOAD_DIR/$EMACS_GZ ];then
    # -O 指定目标文件
    # -x 强制创建目录
    wget -O $DOWNLOAD_DIR/$EMACS_GZ -x $EMACS_DOWNLOAD_URL
fi

cd $DOWNLOAD_DIR

if [ ! -e $EMACS_DIR ];then
    tar zxvf $EMACS_GZ
fi

cd $EMACS_DIR
./configure --prefix=$INSTALL_DIR --with-x-toolkit=no --with-xpm=no --with-jpeg=no --with-gif=no --with-tiff=no --with-gnutls=no
make
make install


