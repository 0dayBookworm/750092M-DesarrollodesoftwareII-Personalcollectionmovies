<div class="page-header">
    <center>
        <h2><strong>DATOS DE USUARIO</strong></h2>
    </center>
</div>
<!--Datos de usuario -->
<form class="form-horizontal " id="ProfileForm" name="ProfileForm" data-toggle="bootstrapValidator">
    <!-- Username and mail-->
    <div class="form-group">
        <label class="control-label col-md-3">Nombre de usuario:</label>
        <div class="col-md-9">
            <input type="text" id="Username" name="Username" class="form-control text-lowercase" value={{.Username}} readonly>
        </div>
    </div>
    <div class="form-group">
        <label class="control-label col-md-3">Email:</label>
        <div class="col-md-9">
            <input type="email" id="Email" name="Email" class="form-control text-lowercase" value={{.Email}} readonly>
        </div>
    </div>
    <!-- Fin username and mail -->
    <!-- Primer nombre -->
    <div class="form-group">
        <label class="control-label col-md-3">Primer nombre:</label>
        <div class="col-md-9">
            <input type="text" if="FirstName" name="FirstName" class="form-control" value={{.FirstName}}>
        </div>
    </div>
    <!-- Segundo nombre -->
    <div class="form-group">
        <label class="control-label col-md-3">Segundo nombre:</label>
        <div class="col-md-9">
            <input type="text" id="SecondName" name="SecondName" class="form-control" value={{.SecondName}}>
        </div>
    </div>
    <!-- Apellidos -->
    <div class="form-group">
        <label class="control-label col-md-3">Apellidos:</label>
        <div class="col-md-9">
            <input type="text" id="LastName" name="LastName" class="form-control" value={{.LastName}}>
        </div>
    </div>
    <!-- Fecha de nacimiento -->
    <div class="form-group">
        <label class="control-label col-md-3">F. Nacimiento:</label>
        <div class="col-md-9">
            <div class="input-group date" id="Date">
                <span class="input-group-addon">
                            <span class="glyphicon glyphicon-calendar"></span>
                </span>
                <input type="text" id="BirthDate" name="BirthDate" class="form-control" value={{date .BirthDate "Y-m-d"}}>
            </div>
        </div>
    </div>
    <!-- Genero -->
    <div class="form-group">
        <label class="control-label col-md-3">Genero:</label>
        <div class="col-md-9">
            <div class="col-md-9">
                <label class="radio-inline">
                        <input type="radio"  id="Gender" name="Gender"  value="male" {{.Male}}> Masculino
                    </label>
            </div>
            <div class="col-md-9">
                <label class="radio-inline">
                        <input type="radio"  id="Gender" name="Gender"  value="female" {{.Female}}> Femenino
                    </label>
            </div>
        </div>
    </div>
    <!-- Fin genero -->
    <!-- Guardar cambios -->
    <br>
    <div class="form-group">
        <div class="col-md-offset-3 col-md-9">
            <button type="submit" id="UpdateProfile" class="btn btn-primary" aria-label="Left Align" disabled>
                    <i class="glyphicon glyphicon-pencil"></i>
                    Guardar Cambios
                </button>
        </div>
    </div>
</form>