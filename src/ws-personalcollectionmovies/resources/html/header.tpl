{{define "header"}}
<!DOCTYPE html>
<html lang="es">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width; initial-scale=1.0">
    <title>{{.Title}} - PersonalCollectionMovies [WSPCM]</title>
    <!-- Estilos CSS vinculados -->
    <link href="//getbootstrap.com/dist/css/bootstrap.css" rel="stylesheet">
    <link href="https://netdna.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" rel='stylesheet' type='text/css'>
    <link href="//maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.5.0/css/bootstrap-datepicker.min.css" />
    <link href="//oss.maxcdn.com/jquery.bootstrapvalidator/0.5.2/css/bootstrapValidator.min.css" rel="stylesheet" />
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-rating/1.3.2/bootstrap-rating.css" rel="stylesheet" />
    <link href="http://cdn.datatables.net/plug-ins/28e7751dbec/integration/bootstrap/3/dataTables.bootstrap.css" rel="stylesheet" />
    <link href="//cdn.datatables.net/1.10.12/css/jquery.dataTables.min.css" rel="stylesheet" />
    <!-- Fonts -->
    <link href='https://fonts.googleapis.com/css?family=Yanone+Kaffeesatz:400,700,300,200&subset=latin,latin-ext' rel='stylesheet' type='text/css'>
    <!-- Estilos propios -->
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/footer.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/footer-basic-centered.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/header.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/login.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/general.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/dropdown.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/panel-horizontal.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/meta-date.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/movies.css" rel="stylesheet">
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/sidebar.css" rel="stylesheet">
    <!-- customJPages css -->
    <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/jPages/customJPages.css" rel="stylesheet">
</head>

<body onload="checkCookies()">
    <!--
    <section id="Preloader">
		<div class="loading-circle fa-spin"></div>
	</section>
	-->
    <!-- Barra de navegación estatica-->
    <!-- Debe contener inicio de seión y registro -->
    <nav class="navbar navbar-fixed-top navbar-inverse">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#NavbarCollapse" aria-expanded="true">
                        <i class="glyphicon glyphicon-menu-hamburger"></i>
                    </button>
                <a href="//personalcollectionmovies-alobaton.c9users.io/" class="navbar-brand">
                    <strong>PersonalCollectionMovies</strong>
                </a>
            </div>
            <div id="NavbarCollapse" name="NavbarCollapse" class="collapse navbar-collapse">
                <!-- Definimos una variable que nos permita cambiar los controles de inicio y cierre de sesión -->
                {{.SessionControl}}
            </div>
        </div>
    </nav>
    {{template "login" .}}
    <!-- Fin barra de navegacion -->
    <!-- Carrusel -->
    <header>
        <div id="Carousel" class="carousel slide" data-ride="carousel">
            <!-- Indicators -->
            <ol class="carousel-indicators">
                <li data-target="#Carousel" data-slide-to="0" class="active"></li>
                <li data-target="#Carousel" data-slide-to="1"></li>
            </ol>
            <!-- Wrapper for slides -->
            <div class="carousel-inner" role="listbox">
                <div class="item active">
                    <img src="//personalcollectionmovies-alobaton.c9users.io/public/images/carrousel_img_1.jpg">
                    <div class="carousel-caption">
                    </div>
                </div>
                <div class="item">
                    <img src="//personalcollectionmovies-alobaton.c9users.io/public/images/carrousel_img_2.jpg">
                    <div class="carousel-caption">
                    </div>
                </div>
            </div>
            <!-- Controls -->
            <a class="left carousel-control" href="#Carousel" role="button" data-slide="prev">
                <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
                <span class="sr-only">Previous</span>
            </a>
            <a class="right carousel-control" href="#Carousel" role="button" data-slide="next">
                <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
                <span class="sr-only">Next</span>
            </a>
        </div>
    </header>
    <!-- Carrusel -->
    <!-- Barra de navegación de pagínas -->
    <nav class="navbar navbar-inverse navbar">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbarMainMenu" aria-expanded="true">
                            <i class="glyphicon glyphicon-menu-hamburger"></i>
                        </button>
            </div>
            <div id="navbarMainMenu" class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li><a href="//personalcollectionmovies-alobaton.c9users.io/">INICIO</a></li>
                    <li><a href="//personalcollectionmovies-alobaton.c9users.io/movie/nowplaying?Page=1">CARTELERA</a></li>
                    <li><a href="//personalcollectionmovies-alobaton.c9users.io/movie/upcoming?Page=1">ESTRENOS</a></li>
                    <li><a href="//personalcollectionmovies-alobaton.c9users.io/movie/popular?Page=1">POPULARES</a></li>
                    <li><a href="#">CONOCENOS</a></li>
                    <li><a href="//personalcollectionmovies-alobaton.c9users.io/contact">CONTACTENOS</a></li>
                </ul>
            </div>
        </div>
    </nav>
    <!-- Fin barra de navegacion paginas -->
    <!-- Buscar pelicula -->
    <div class="container">
        <form class="form-horizontal" id="SearchForm" name="SearchForm">
            <div class="input-group">
                <span class="input-group-btn">
                        <button class="btn btn-primary" type="button" id="Search" name="Search">
                           <i class="glyphicon glyphicon-search"></i>
                           BUSCAR PELíCULA
                        </button>
                    </span>
                <input type="text" name="Title" class="form-control" placeholder="Buscar...">
            </div>
        </form>
    </div>
    <!-- Buscar pelicula modal -->
    {{template "search" .}}
    <!-- Fin header -->
    {{end}}