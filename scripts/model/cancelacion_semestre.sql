-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.8.1
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---

-- object: cancelacion_semestre | type: ROLE --
-- DROP ROLE IF EXISTS cancelacion_semestre;
CREATE ROLE cancelacion_semestre WITH 
	ENCRYPTED PASSWORD 'sistemasoas';
-- ddl-end --
COMMENT ON ROLE cancelacion_semestre IS 'El usuario usado para las transacciones de cancelacion de semestre';
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: new_database | type: DATABASE --
-- -- DROP DATABASE IF EXISTS new_database;
-- CREATE DATABASE new_database
-- ;
-- -- ddl-end --
-- 

-- object: academica_estudiantes | type: SCHEMA --
-- DROP SCHEMA IF EXISTS academica_estudiantes CASCADE;
CREATE SCHEMA academica_estudiantes;
-- ddl-end --
ALTER SCHEMA academica_estudiantes OWNER TO cancelacion_semestre;
-- ddl-end --

SET search_path TO pg_catalog,public,academica_estudiantes;
-- ddl-end --

-- object: academica_estudiantes.tipo_cancelacion_semestre | type: TABLE --
-- DROP TABLE IF EXISTS academica_estudiantes.tipo_cancelacion_semestre CASCADE;
CREATE TABLE academica_estudiantes.tipo_cancelacion_semestre(
	id serial NOT NULL,
	tipo character varying(50) NOT NULL,
	descripcion character varying(200),
	CONSTRAINT tipo_cancelacion_semestre_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica_estudiantes.tipo_cancelacion_semestre IS 'especifica el tipo de cancelacion, retiro voluntario o no renovacion de matricula(aplazar semestre)';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.tipo_cancelacion_semestre.id IS 'una secuencia que identifica el tipo de cancelacion';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.tipo_cancelacion_semestre.tipo IS 'el nombre del tipo de descripcion, corto y significativo';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.tipo_cancelacion_semestre.descripcion IS 'una descripcion profunda sobre el tipo registrado';
-- ddl-end --
ALTER TABLE academica_estudiantes.tipo_cancelacion_semestre OWNER TO cancelacion_semestre;
-- ddl-end --

-- object: academica_estudiantes.estado_cancelacion_semestre | type: TABLE --
-- DROP TABLE IF EXISTS academica_estudiantes.estado_cancelacion_semestre CASCADE;
CREATE TABLE academica_estudiantes.estado_cancelacion_semestre(
	id serial NOT NULL,
	estado character varying(50) NOT NULL,
	descripcion character varying(200),
	CONSTRAINT estado_cancelacion_semestre_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.estado_cancelacion_semestre.id IS 'identificador serial de estado_cancelacion_semestre';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.estado_cancelacion_semestre.estado IS 'Nombre corto del estado de la cancelacion';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.estado_cancelacion_semestre.descripcion IS 'descripcion detallada del significado del estado de la cancelacion';
-- ddl-end --
ALTER TABLE academica_estudiantes.estado_cancelacion_semestre OWNER TO cancelacion_semestre;
-- ddl-end --

-- object: academica_estudiantes.cancelacion_semestre | type: TABLE --
-- DROP TABLE IF EXISTS academica_estudiantes.cancelacion_semestre CASCADE;
CREATE TABLE academica_estudiantes.cancelacion_semestre(
	id serial NOT NULL,
	id_tipo integer NOT NULL,
	id_estado integer NOT NULL,
	id_estudiante bigint,
	uid integer,
	motivo character varying(2000),
	observaciones character varying(2000),
	num_folios_anexados smallint,
	motivo_estado_aprobado character varying(2000),
	fecha_solicitud timestamp,
	fecha_aprobacion timestamp,
	CONSTRAINT cancelacion_semestre_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica_estudiantes.cancelacion_semestre IS 'Se consagran los datos relacionados al evento de cancelar semestre';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.id IS 'identificador serial del evento';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.id_estado IS 'identificador de la tabla estado_cancelacion_semestre';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.id_estudiante IS 'es el codigo del estudiante que solicita la cancelacion';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.uid IS 'Identificador del sistema tuleap';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.motivo IS 'motivo por el cual se aplaza, cancela semestre, este es rellenado por el estudiante';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.observaciones IS 'observaciones adicionales al proceso, se llena por el estudiante';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.num_folios_anexados IS 'numero de folios anexados, aparece en el formato de cancelacion';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.motivo_estado_aprobado IS 'el motivo por el cual se aprueba o deniega la solicitud, es llenado por el funcionario correspondiente';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.fecha_solicitud IS 'fecha en que se realiza la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica_estudiantes.cancelacion_semestre.fecha_aprobacion IS 'fecha en que se aprueba la solicitud';
-- ddl-end --
ALTER TABLE academica_estudiantes.cancelacion_semestre OWNER TO cancelacion_semestre;
-- ddl-end --

-- object: tipo_cancelacion_semestre_id_tipo_fkey | type: CONSTRAINT --
-- ALTER TABLE academica_estudiantes.cancelacion_semestre DROP CONSTRAINT IF EXISTS tipo_cancelacion_semestre_id_tipo_fkey CASCADE;
ALTER TABLE academica_estudiantes.cancelacion_semestre ADD CONSTRAINT tipo_cancelacion_semestre_id_tipo_fkey FOREIGN KEY (id_tipo)
REFERENCES academica_estudiantes.tipo_cancelacion_semestre (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: estado_cancelacion_semestre_id_estado_fkey | type: CONSTRAINT --
-- ALTER TABLE academica_estudiantes.cancelacion_semestre DROP CONSTRAINT IF EXISTS estado_cancelacion_semestre_id_estado_fkey CASCADE;
ALTER TABLE academica_estudiantes.cancelacion_semestre ADD CONSTRAINT estado_cancelacion_semestre_id_estado_fkey FOREIGN KEY (id_estado)
REFERENCES academica_estudiantes.estado_cancelacion_semestre (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


