package main


import(
	// CORE LIBS
	rn "math/rand"
	"time"
	"log"
	"os"
	"fmt"
	"net/http"

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


	// SERVER
	app := http.Server{
		Handler:	routes,
		Addr:		host+":"+port,

		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// SIMPLE WAY TO START SERVER
	log.Fatal(app.ListenAndServe())

	// TODO COMPLEX WAY TO START AND MANAGE SERVER
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	
	// go func(){
	// 	defer close(sigs)
	// 	if serr := app.ListenAndServe(); serr != http.ErrServerClosed{
	// 		switch err{
	// 		case nil:
	// 			err = fmt.Errorf("unable to start server: %v", serr)
	// 		default:
	// 			err = errors.Wrapf(err, "unable to start server: %v", serr)

	// 		}
	// 	}

	// }()

	// log.Printf("Server sterted succesfully")
	// s := <-sigs
	// err = fmt.Errorf("shutting down server because of: %v: %v", s, err)
	
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// stx, sc := context.WithTimeout(ctx, time.Second*5)
	// defer sc()
	// serr := app.Shutdown(stx)
	// switch serr {
	// case nil:
	// 	log.Printf("server stopped gracefully")
	// default:
	// 	log.Printf("unable to stop server: %v", serr)
	// }





}