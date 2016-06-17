{{template "header" .}}
<div class="container">
    <center>
        <h1>    
            <strong>
               404 {{.ErrorCode}}
            </strong>
        </h1>
        <h2>
            Prueba {{.ErrorReason}}
        </h2>
        <p align="center">
            {{.ErrorMessage}}
        </p>
    </center>
</div>
{{template "footer" .}}