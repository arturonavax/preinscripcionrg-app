API-GraphQL del Sistema de Pre-incripcion del Romulo Gallegos.

Servidor API-GraphQL de un sistema de Pre-inscripcion de una institucion educativa, Con el cual los representantes puedan pre-inscribir sus alumnos en la institucion e imprimir la planilla ya con todos los datos adjuntados.

La institucion podra ver todos los alumnos pre-inscritos en la base de datos manipulada por este servidor, Realizar seguimientos de todos los datos de los alumnos pre-inscritos y automatizar el proceso de realizar todo esto a papel.

Este servidor es totalmente compatible con cualquier cliente Frontend que tenga la capacidad de enviar los formatos JSON y GraphQL al servidor, La base de datos implementada es PostgreSQL, Totalmente alternable entre PostgreSQL, MySQL, SQLite3 y SQL Server ya que utiliza un ORM llamado GORM.

Se decidio la implementacion del lenguaje de programacion Go por su versatilidad como servidor y altos niveles de concurrencia que permiten un rendimiento superior a la mayoria de lenguajes para servidores.

La autenticacion por JSON Web Token se implementa por la velocidad, Facilidad de uso y seguridad, Nos permite hacer un API mas escalable.

-------------------------------------------

La estructura de carpetas para el funcionamientos de este servidor en formato ejecutable es la siguiente :

    preinscripcion  (Ejecutable)
    databases/
        configDB.json   (Archivo de configuracion de la base de datos)
    
    security/
        keys/   (Carpetas donde deben estar las llaves de encriptacion para generar JWT)
            private.rsa
            public.rsa

-------------------------------------------

Banderas validas para pasar al ejecutable :
    preinscripcion --port 9090  (Cambiar el puerto por defecto del servidor (8080)).
    preinscripcion --createTables   (Crea todas las tablas necesarias en la base de datos).
    preinscripcion --dropTables (Elimina todas las tablas de la base de datos).
    preinscripcion --recreateTables (Elimina y vuelve a crear todas las tablas necesarias en la base de datos).
    preinscripcion --initProduction (Crea todas las carpetas y archivos necesarios para ejecutar la aplicacion).

    NOTA : Al ejecutar el servidor, Directamente crea todas las carpetas y archivos que necesita para funcionar.

-------------------------------------------

Rutas de la API :

    /api/register   -   A esta ruta se le envia por el metodo POST un JSON con la siguiente estructura :
        HeaderKey: Content-Type  ;  Value: application/json

        {
            "firstName" : "Arturo",
            "lastName" : "Nava",
            "email" : "arthurnavah@gmail.com",
            "address" : "Valle Frio",
            "phoneNumber" : "+584166651924",
            "username" : "ArthurNavaH",
            "password" : "1010",
            "confirmPassword" : "1010"
        }

    ----------------------

    /api/login  -   A esta ruta se le envia por el metodo POST un JSON con la siguiente estructura :
        HeaderKey: Content-Type  ;  Value: application/json

        {
            "username": "ArthurNavaH",
            "passowrd: "1010"
        }

    Devuelve un JSON Web Token que es con el cual podremos acceder a la mayor parte de funcionalidades de el sistema.
    ----------------------

    /graphql    -   A esta ruta se le envia por el metodo POST un paquete GraphQL con la siguiente estructura :
        HeaderKey: Content-Type  ;  Value: application/graphql
        HeaderKey: Authorization  ;  Value: (JWT Devuelto por /api/login)

        mutation {
            studentC (
                countryOfBirthID : 1,
                stateOfBirthID : 1,
                municOfBirthID : 1,
                municipalityID : 1,
                institutionOfOriginID : 1,
                parishID : 1,
                sectorID : 1,
                typeOfRoadID : 1,
                fatherID : 1,
                motherID : 1,
                mentionID : 1,
                sectionID : 1,
                studentConditionID : 1,
                representativeID : 1,
                teacherID : 1,
                
                SIGECOD : "A101",
                firstName : "Frank",
                lastName : "Peñaranda",
                ciType : "V",
                ci : 29100800,
                dateOfBirth : "11/11/01",
                gender : "M",
                healthProblem : false,
                healthProblemE : "",
                email : "coffemanfp@gmail.com",
                phoneNumber : "+584167634291",
                address : "Valle Frio",
                scholarship : true,
                canaima : true,
                conditionOfHousing : "PROPIA",
                year : 4,
                age : 15,
                size : "L",
                weight : 60,
                repeatAsignature : "Castellano",
                pendingAsignature : "",
                regular : true,
                inscriptionDate : "10/08/15"
                
            ){id,message,code}
        }

    ----------------------
-------------------------------------------


Paquetes utilizados para realizar este proyecto :
    jwt-go  -   Para la manipulacion de JSON Web Tokens.
    graphql-go/graphql  -   Para la estructuracion del API GraphQL.
    graphql-go/handler  -   Para diseñar el Handler que pueda manipular GraphQL de manera correcta.
    gorm    -   Para la manipulacion de SQL por medio de un ORM.
    gorm/dialects/postgres" -   Como el Driver de PostgreSQL.