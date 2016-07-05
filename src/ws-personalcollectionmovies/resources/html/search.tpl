{{define "search"}}
<div class="modal fade" id="SearchModal" name="SearchModal" tabindex="-1" role="dialog" aria-labelledby="SearchModalTitle" aria-hidden="true">
    <div class="modal-dialog ">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="SearchModalTitle">BUSCAR PELÍCULA</h4>
            </div>
            <div class="modal-body">
                <!-- Mostramos la lista de peliculas -->
                <ul class="list-group" id="SearchResult" name="SearchResult"  style="width: 100%; height: 400px; overflow: auto">
                    
                </ul>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">Cerrar</button>
            </div>
        </div>
    </div>
</div>
{{ end }}