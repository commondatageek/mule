variable ami {
    type = string
    default = "ami-078278691222aee06" # ubuntu 20.04, arm64; 
    nullable = false
    # default = "ami-074251216af698218" # for ubuntu 20.04, amd64
}

variable instance_type {
    type = string
    default = "t4g.nano" # is an arm64 instance (graviton2 processor)
    nullable = false
}

variable subnet_id {
    type = string
    nullable = false
}

variable vpc_security_group_ids {
    type = list(string)
    nullable = false
}

variable key_name {
    type = string
    nullable = false
}

variable name {
    type = string
    default = "skymule"
}