/^ *!/ {
  sub(/^ *!/, "", $0)
  s = s " " $0
}
/begin_src go/ {
  if(s) {
    gsub(/\\\[^{]+{/, "", s)
    gsub(/}/, "", s)
    gsub(/\$/, "", s)
    printf "%s\\n//%s\\n", $0, s
    s = ""
  } else 
    print
}
!/^ *!/ && !/begin_src (go|c|c++|r|python|sh|html)/ {
  print
}
