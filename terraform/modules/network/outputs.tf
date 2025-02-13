output "private_subnet_a" {
  value = {
    id                = aws_subnet.private_a.id
    cidr_block        = aws_subnet.private_a.cidr_block
    availability_zone = aws_subnet.private_a.availability_zone
    arn               = aws_subnet.private_a.arn
  }
  description = "Details of the private subnet"
}

output "private_subnet_b" {
  value = {
    id                = aws_subnet.private_b.id
    cidr_block        = aws_subnet.private_b.cidr_block
    availability_zone = aws_subnet.private_b.availability_zone
    arn               = aws_subnet.private_b.arn
  }
  description = "Details of the private subnet"
}

output "public_subnet" {
  value = {
    id                = aws_subnet.public.id
    cidr_block        = aws_subnet.public.cidr_block
    availability_zone = aws_subnet.public.availability_zone
    arn               = aws_subnet.public.arn
  }
  description = "Details of the private subnet"
}
