#!/bin/bash
cd /home/pi/workdir/hexo/jinxycandotailwhip
git checkout /home/pi/workdir/hexo/jinxycandotailwhip/source/_posts/*
git pull -f
hexo generate --cwd /home/pi/workdir/hexo/jinxycandotailwhip
