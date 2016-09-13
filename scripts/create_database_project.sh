#!/bin/bash
newgrp docker
docker run -d -p 5432:5432 -e POSTGRESQL_USER=cancelacion_semestre -e POSTGRESQL_PASS=docker -e POSTGRESQL_DB=cancelacion_semestre orchardup/postgresql
psql -h localhost -d udistrital -U cancelacion_semestre -W < model/cancelacion_semestre.sql
#bee api github.com/JorgeUlises/oas_be_cancelacion_semestre -driver=postgres -conn=postgres://cancelacion_semestre:docker@localhost:5432/udistrital?sslmode=disable&search_path=cancelacion_semestre
#bee run -downdoc=true -gendoc=true
