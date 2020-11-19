package main


import(
	// CORE LIBS
	rn "math/rand"
	"time"
	"log"
	"os"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"context"

	// EXTERNAL LIBS
	"github.com/joho/godotenv"

	// INTERNAL PACKAGES
	"go-rest-starter/interfaces"


)

// SET UP CUMSTOM LOGGING
func init(){
	rn.Seed(time.Now().UnixNano())

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main(){
	// INITIALIZE THE ERROR OBJECT
	var err error
	// RUN AFTER MAIN IS EXITED  
	defer func(){
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// LOAD ENVIRONMENT VARIABLES
	if logerr := godotenv.Load(); logerr != nil {
		err = fmt.Errorf("env file is missing")
		return
	}

	// LOAD HTTP VARIABLES
	port := os.Getenv("API_PORT")
	host := os.Getenv("HOST")
	if port == "" {
		err = fmt.Errorf("missing API_PORT")
		return
	}
	if host == "" {
		err = fmt.Errorf("missing HOST")
		return
	}
	routes := interfaces.Routing()


	// CREATE APPLICATION CONTEXT
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	// SERVER
	app := http.Server{
		Handler:	routes,
		Addr:		host+":"+port,

		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}


	// CREATE A SYSCALL CHANNEL TO MAINTAIN CONNECTION
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	
	go func(){
		defer close(sigs)
		if serr := app.ListenAndServe(); serr != http.ErrServerClosed{
			err = fmt.Errorf("Unable to start server: %v", serr)
		} 

	}()
	log.Printf("Server started succesfully")
	s := <-sigs
	err = fmt.Errorf("Shutting down server because of: %v: %v", s, err)

	// SHUTDOWN SERVER PROCESSES  
	stx, sc := context.WithTimeout(ctx, time.Second*5)
	defer sc()
	serr := app.Shutdown(stx)
	switch serr {
	case nil:
		log.Printf("server stopped gracefully")
	default:
		log.Printf("unable to stop server: %v", serr)
	}


}