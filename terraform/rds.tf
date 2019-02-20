resource "aws_db_instance" "app-prod" {
    identifier                = "app-prod"
    allocated_storage         = 200
    storage_type              = "gp2"
    engine                    = "mysql"
    engine_version            = "5.7.17"
    instance_class            = "db.t2.large"
    apply_immediately         = true
}
