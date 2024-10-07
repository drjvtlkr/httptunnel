# SSH Tunneling to HTTP #

1. This program sends the file to the HTTP directly

2. Clone the repo by copying the command
    ```bash
    git clone https://github.com/drjvtlkr/httptunnel.git

3. Now run the program on the vs code terminal
    ```bash 
    go run main.go

4. Now open a sister terminal in the same project folder where there is main.go 

5. Now run the command 
    ```bash
    ssh localhost -p 2222 < main.go

6. You will get an warning initially, which says
    ```bash
    Pseudo-terminal will not be allocated because stdin is not a terminal.
    The authenticity of host '[localhost]:2222 ([127.0.0.1]:2222)' can't be established.
    RSA key fingerprint is SHA256:<YOUR_UNIQUE_RSA_KEY>
    This key is not known by any other names.

7. Press Yes and continue    

8. What this does is sends the main.go file onto the location in the browser
    ```bash
    http://localhost:3000/?id=<your_tunnel_id>

9. You will get this tunnel ID in your terminal where you had run 

    ```bash
    go run main.go

10. Voila, you are done!!

## Happy coding