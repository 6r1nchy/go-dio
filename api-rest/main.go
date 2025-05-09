package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
    "strconv"
)

// Estrutura de Endereço
type Endereco struct {
    Rua        string `json:"rua"`
    Numero     int    `json:"numero"`
    Complemento string `json:"complemento"`
    Cidade     string `json:"cidade"`
    Estado     string `json:"estado"`
}

// Estrutura do Cliente
type Cliente struct {
    ID        int      `json:"id"`
    Nome      string   `json:"nome"`
    Sobrenome string   `json:"sobrenome"`
    Telefone  string   `json:"telefone"`
    Endereco  Endereco `json:"endereco"`
}

// "Banco de dados fake" utilizando um slice para armazenar os clientes
var clientes []Cliente

// Criar um novo cliente
func criarCliente(w http.ResponseWriter, r *http.Request) {
    var cliente Cliente
    json.NewDecoder(r.Body).Decode(&cliente)
    cliente.ID = rand.Intn(1000) // Gera um ID aleatório
    clientes = append(clientes, cliente)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cliente)
}

// Listar todos os clientes
func listarClientes(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(clientes)
}

// Buscar um único cliente pelo ID
func buscarCliente(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, cliente := range clientes {
        if cliente.ID == id {
            json.NewEncoder(w).Encode(cliente)
            return
        }
    }
    http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

// Excluir um cliente pelo ID
func excluirCliente(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, cliente := range clientes {
        if cliente.ID == id {
            clientes = append(clientes[:index], clientes[index+1:]...)
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "Cliente com ID %d excluído com sucesso", id)
            return
        }
    }
    http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

// Editar um cliente pelo ID
func editarCliente(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    for index, cliente := range clientes {
        if cliente.ID == id {
            var novoCliente Cliente
            json.NewDecoder(r.Body).Decode(&novoCliente)
            novoCliente.ID = id // Mantém o mesmo ID
            clientes[index] = novoCliente

            json.NewEncoder(w).Encode(novoCliente)
            return
        }
    }
    http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

// Função principal
func main() {
    r := mux.NewRouter()

    r.HandleFunc("/clientes", criarCliente).Methods("POST")
    r.HandleFunc("/clientes", listarClientes).Methods("GET")
    r.HandleFunc("/clientes/{id}", buscarCliente).Methods("GET")
    r.HandleFunc("/clientes/{id}", excluirCliente).Methods("DELETE")
    r.HandleFunc("/clientes/{id}", editarCliente).Methods("PUT")

    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", r)
}
