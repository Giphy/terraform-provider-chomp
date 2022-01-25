terraform {
  required_providers {
    chomp = {
      version = ">= 0.1"
      source  = "giphy/chomp"
    }
  }
}

variable "hosted_zones" {
  type = map(string)
  default = {
    "example.com"         = "FG2923A8E"
    "example.org"         = "EDZF29B43"
    "sub.example.org"     = "7AB23ECF2"
    "alt.sub.example.org" = "4DCB33FB8"
  }
}

data "chomp_left" "zone" {
  lookup    = "other.zone.sub.example.org"
  separator = "."

  src = var.hosted_zones
}

locals {
    key = data.chomp_left.zone.key
    val = var.hosted_zones[local.key]
}

# Returns hosted zone key
output "zone_key" {
  value = local.key
}

# Returns hosted zone value
output "zone_value" {
  value = local.val
}

