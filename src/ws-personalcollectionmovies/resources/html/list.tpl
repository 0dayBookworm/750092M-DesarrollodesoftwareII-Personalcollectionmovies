<div class="container">
    <!-- Page header -->
    <div class="page-header">
        <center>
            <h1><strong>{{.Title}}</strong></h1>
        </center>
    </div>
    <!-- Las peliculas de la busqueda -->
    <ul id="ViewListByCinema" class="list-group">
        {{str2html .ListItems}}
    </ul>
    <!-- Paginador -->
    <div class="row">
        <div class="col-xs-1">
            <p>({{.Total}} resultados)</p>
        </div>
        <div class="pull-right">
            <ul class="pagination"></ul>
        </div>
    </div>
</div>
