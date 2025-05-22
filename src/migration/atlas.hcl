data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./",
  ]
}

locals {
  mysql_url = "mysql://root:secret@localhost:3306/trading"
  pg_url = "postgres://postgres:postgres@localhost:5432/trading?sslmode=disable"
}

env "mysql" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  url = local.mysql_url
  migration {
    dir = "file://migrations/mysql"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "postgres" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  url = local.pg_url
  migration {
    dir = "file://migrations/postgres"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
