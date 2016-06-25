{{template "header" .}}

<div class="container">
    <!-- Page header -->
    <div class="page-header">
        <center>
            <h1><strong>MANTENTE EN CONTACTO</strong></h1>
        </center>
    </div>
    <!-- Contact form -->
    <form class="form-horizontal">
        <div class="form-group">
            <label for="InputName" class="control-label col-xs-2">Nombre:</label>
            <div class="col-xs-10">
                <input type="name" class="form-control" placeholder="Nombre">
            </div>
        </div>
        <div class="form-group">
            <label for="InputEmail" class="control-label col-xs-2">Email:</label>
            <div class="col-xs-10">
                <input type="email" class="form-control" placeholder="Email">
            </div>
        </div>
        <div class="form-group">
            <label for="InputPassword" class="control-label col-xs-2">Asunto:</label>
            <div class="col-xs-10">
                <input type="text" class="form-control" placeholder="Asunto">
            </div>
        </div>
        <div class="form-group">
            <label for="InputMessage" class="control-label col-xs-2">Asunto:</label>
            <div class="col-xs-10">
                <textarea rows="4" placeholder="Escribe tu mensaje" required="required" class="form-control"></textarea>
            </div>
        </div>
        <div class="form-group">
            <div class="col-xs-offset-2 col-xs-10">
                <button type="submit" class="btn btn-primary">
                    <i class="glyphicon glyphicon-ok"></i>
                    Enviar
                </button>
            </div>
        </div>
    </form>
</div>

{{template "footer" .}}