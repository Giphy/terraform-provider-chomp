data "chomp_left" "zone" {
  lookup    = "my.sub.example.org"
  separator = "."

  src = {
    "sub.example.org" = "DEADBEEF"
    "example.org"     = "BEEFFEED"
  }
}

