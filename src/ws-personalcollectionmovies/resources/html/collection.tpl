{{template "header" .}}

<div class="container">
    <div class="page-header">
        <center>
            <h1><strong>MIS COLECCIONES</strong></h1>
        </center>
    </div>
    <table id="Collections" name="Collections" class="table table-inverse table-sm table-striped" cellspacing="0" width="100%" >
        <thead>
            <tr>
                <th>#</th>
                <th>Nombre de la lista</th>
                <th>Descripción</th>
                <th>Editar</th>
                <th>Eliminar</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <th scope="row">1</th>
                <td>Vistas</td>
                <td>Peliculas vistas</td>
                <td>
                    
                </td>
                <td>
                    
                </td>
            </tr>
            <tr>
                <th scope="row">2</th>
                <td>Por ver</td>
                <td>Peliculas por ver</td>
                <td>
                    
                </td>
                <td>
                    
                </td>
            </tr>
            <tr>
                <th scope="row">3</th>
                <td>Favoritas</td>
                <td>Mi lista de favoritas</td>
                <td>
                    <a href="#">
                        <span class="glyphicon glyphicon-pencil"></span>
                    </a>
                </td>
                <td>
                    <a href="#">
                        <i class="glyphicon glyphicon-trash"></i>
                    </a>
                </td>
            </tr>
            <tr>
                <th scope="row">4</th>
                <td>Suspenso</td>
                <td>Peliculas de suspenso</td>
                <td>
                    <a href="#">
                        <span class="glyphicon glyphicon-pencil"></span>
                    </a>
                </td>
                <td>
                    <a href="#">
                        <i class="glyphicon glyphicon-trash"></i>
                    </a>
                </td>
            </tr>
        </tbody>
    </table>
</div>

{{template "footer" .}}