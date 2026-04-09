#!/bin/bash
sed '
s/^#+begin_src *latex/@/
s/^#+begin_export *latex/@/
s/^#+begin_src *[cC] *<</<</
s/^#+begin_src *[cC]++ *<</<</
s/^#+begin_src *[sS][hH] *<</<</
s/^#+begin_src *[aA][wW][kK] *<</<</
s/^#+begin_src *[gG][oO] *<</<</
s/^#+begin_src *[hH][tT][mM][lL] *<</<</
s/^#+begin_src *[pP][yY][tT][hH][oO][nN] *<</<</
s/^#+begin_src *[rR] *<</<</

s/\/\/ *<</<</
/^#+end/d
/^\*/d
s/^  //
' $@
