{{define "places"}}
<div class="modal fade" id="PlacesModal" name="PlacesModal" tabindex="-1" role="dialog" aria-labelledby="PlacesModalTitle" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="PlacesModalTitle">SELECCONA UNA UBICACIÓN</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal" id="PlacesForm" name="PlacesForm">
                    <div class="input-group">
                        <span class="input-group-btn">
                                <button class="btn btn-primary" type="button" id="Places" name="Places">
                                   <i class="glyphicon glyphicon-search"></i>
                                   BUSCAR
                                </button>
                            </span>
                        <input type="text" name="Place" class="form-control" placeholder="¿Donde viste esta película?">
                    </div>
                </form>
                <!-- Mostramos la lista de peliculas -->
                <br>
                <ul class="list-group" id="PlacesResult" name="PlacesResult" style="width: 100%; height: 400px; overflow: auto">
                    
                </ul>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">Cerrar</button>
            </div>
        </div>
    </div>
</div>
{{ end }}