\c preinscripcion

-- CREACION DE PERSONAS.
INSERT INTO peoples (first_name, last_name, email, phone_number, address) VALUES
	('Arturo', 'Nava', 'arthurnavah@gmail.com','+584166651924', 'Valle Frio'),
	('Maikol', 'Hernandez', 'maikol@gmail.com', '04146360146', 'Chinita'),
	('Ricardo', 'Romero', 'ricardo@gmail.com', '041664091   70', 'La Victoria'),
	('Carlos', 'Vasquez', 'carlos@gmail.com', '04243081929', 'La Limpia')
;

-- CREACION DE USUARIOS.
INSERT INTO users (people_id,status_id,kind_id,username,password) VALUES
	(1, 1, 1, 'ArthurNavaH', 'a8e8300d94b085dd82528964c2a18810a35b2026be4453b400a2bed3593cec4b'),
	(2, 1, 1, 'PitbullBlood', 'a8e8300d94b085dd82528964c2a18810a35b2026be4453b400a2bed3593cec4b'),
	(3, 1, 1, 'MaikolH', 'a8e8300d94b085dd82528964c2a18810a35b2026be4453b400a2bed3593cec4b'),
	(4, 1, 1, 'CarlosV', 'a8e8300d94b085dd82528964c2a18810a35b2026be4453b400a2bed3593cec4b')
;

-- CREACION DE PAISES.
INSERT INTO countries (user_id,name) VALUES
    (1, 'Venezuela'),
    (1, 'Peru'),
    (1, 'Colombia'),
    (1, 'Ecuador'),
    (1, 'Mexico'),
    (1, 'Brazil'),
    (1, 'Chile'),
    (1, 'Costa Rica')
;

-- CREACION DE ESTADOS.
INSERT INTO states (user_id, country_id,name) VALUES
    (1, 1,'Amazonas'),
    (1, 1,'Anzoategui'),
    (1, 1,'Apure'),
    (1, 1,'Aragua'),
    (1, 1,'Barinas'),
    (1, 1,'Bolivar'),
    (1, 1,'Carabobo'),
    (1, 1,'Cojedes'),
    (1, 1,'Distrito Capital'),
    (1, 1,'Delta Amacuro'),
    (1, 1,'Falcon'),
    (1, 1,'Guarico'),
    (1, 1,'Lara'),
    (1, 1,'Merida'),
    (1, 1,'Miranda'),
    (1, 1,'Monagas'),
    (1, 1,'Nueva Esparta'),
    (1, 1,'Portuguesa'),
    (1, 1,'Sucre'),
    (1, 1,'Tachira'),
    (1, 1,'Trujillo'),
    (1, 1,'Vargas'),
    (1, 1,'Yaracuy'),
    (1, 1,'Zulia')
;

-- CREACION DE MUNICIPIOS.
INSERT INTO municipalities (user_id, state_id,name) VALUES
	(1, 23,'Bolívar'),
	(1, 23,'Baralt'),
	(1, 23,'Cabimas'),
	(1, 23,'Catatumbo'),
	(1, 23,'Colón'),
	(1, 23,'Guajira'),
	(1, 23,'Padilla'),
	(1, 23,'Pulgar'),
	(1, 23,'Lossada'),
	(1, 23,'Semprún'),
	(1, 23,'La Cañada de Urdaneta'),
	(1, 23,'Lagunillas'),
	(1, 23,'Machiques (Machiques)'),
	(1, 23,'Mara'),
	(1, 23,'Maracaibo'),
	(1, 23,'Miranda'),
	(1, 23,'Rosario'),
	(1, 23,'San Francisco'),
	(1, 23,'Santa Rita'),
	(1, 23,'Sucre'),
	(1, 23,'Valmore Rodríguez')
;

-- CREACION DE INSTITUTOS.
INSERT INTO institutions (user_id,name) VALUES
    (1, 'Romulo Gallegos'),
    (2, 'Nestor Luis Perez')
;

-- CREACION DE REPRESENTANTES.
INSERT INTO representatives (user_id, first_name, last_name, email, phone_number, ci, relationship, address) VALUES
    (1, 'Zully', 'Marina', 'zullymarina.matheus1@gmail.com', '04146360146', 'V-8500827', 'Madre', 'Valle Frio')
;

-- CREACION DE PROFESORES.
INSERT INTO teachers (user_id, first_name, last_name, email, phone_number, ci) VALUES
    (1, 'Hosman', 'Salamanca', 'hosman@gmail.com', '04146364823', 'V-8507329')
;

-- CREACION DE PARROQUIAS.
INSERT INTO parishes (user_id, municipality_id, name) VALUES 
    (1, 1, 'Antonio Borjas Romero'),
    (1, 1, 'Bolívar'),
    (1, 1, 'Cacique Mara'),
    (1, 1, 'Carracciolo Parra Pérez'),
    (1, 1, 'Cecilio Acosta'),
    (1, 1, 'Cristo de Aranza'),
    (1, 1, 'Coquivacoa'),
    (1, 1, 'Chiquinquira'),
    (1, 1, 'Francisco Eugenio Bustamante'),
    (1, 1, 'Idelfonzo Vásquez'),
    (1, 1, 'Juana de Ávila'),
    (1, 1, 'Luis Hurtado Higuera'),
    (1, 1, 'Manuel Dagnino'),
    (1, 1, 'Olegario Villalobos'),
    (1, 1, 'Raúl Leoni'),
    (1, 1, 'Santa Lucía'),
    (1, 1, 'Venancio Pulgar'),
    (1, 1, 'San Isidro')
;

-- CREACION DE SECTORES.
INSERT INTO sectors (user_id, parish_id,name) VALUES
    (1, 1, 'Valle Frio')
;

-- CREACION DE TIPOS DE VIAS.
INSERT INTO type_of_roads (user_id, name) VALUES
    (1, 'Autobus')
;

-- CREACION DE PADRES.
INSERT INTO fathers (user_id, first_name, last_name, email, phone_number, ci) VALUES
    (1, 'Arturo', 'Enrique', 'pimosoft@gmail.com', '04146360146', 'V-8501612')    
;

-- CREACION DE MADRES.
INSERT INTO mothers (user_id, first_name, last_name, email, phone_number, ci) VALUES
    (1, 'Zully', 'Marina', 'zully@gmail.com', '04146360146', 'V-8500612')    
;

-- CREACION DE MENCIONES.
INSERT INTO mentions (user_id, name) VALUES
    (1, 'INFORMATICA'),
    (1, 'CONTABILIDAD')
;

-- CREACION DE SECCIONES.
INSERT INTO sections (user_id, name) VALUES
    (1, 'A'),
    (1, 'B'),
    (1, 'C'),
    (1, 'D')
;

-- CREACION DE CONDICIONES DE ESTUDIANTES.
INSERT INTO condition_of_housings (user_id, name) VALUES
    (1, 'PROPIA'),
    (1, 'ALQUILADA'),
    (1, 'AL CUIDO'),
    (1, 'CONSERJERIA')
;

-- CREACION DE ASIGNATURAS.
INSERT INTO asignatures (user_id, name) VALUES
    (1, 'MATEMATICA'),
    (1, 'FISICA'),
    (1, 'QUIMICA'),
    (1, 'LENGUAS EXTRANJERAS'),
    (1, 'CASTELLANO'),
    (1, 'BIOLOGIA'),
    (1, 'EDUC.FISICA'),
    (1, 'INFORMATICA'),
    (1, 'CONTABILIDAD'),
    (1, 'TURISMO'),
    (1, 'MECANOGRAFIA'),
    (1, 'DERECHO'),
    (1, 'ARTISTICA'),
    (1, 'PROGRAMACION'),
    (1, 'MANT.TECNICO'),
    (1, 'PREMILITAR'),
    (1, 'OFIMATICA')
;