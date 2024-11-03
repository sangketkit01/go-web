1.Make your own project structure
2.Good design in render template by create template cache by get all the templates into a map
get a request template from a cache if it's not there , add it , execute a request template into a buffer 
, then write it into the Response writer
3.Create configuration by using struct
4.Don't share those configurations to all paths using Repo struct in Handlers
5.Using a new handler by create a function 
6.Need to create a struct(s) to pass values to the template

Useful Third-party libraries :
    Chi , PAT : Route organization
    Gin (Framework) : A powerful framework
    Nosurf : For CSRF Token
    Alexedwards SCS : For Session


Need to remember this :

    var Repo *Repository
    
    type Repository struct {
        App *config.AppConfig
    }
    
    func NewRepository(app *config.AppConfig) *Repository {
        return &Repository{
            App: app,
        }
    }
    
    func NewHandlers(r *Repository) {
        Repo = r
    }

------------------------------------------------------------------------------------------------

    var app *config.AppConfig
    var functions = template.FuncMap{}
    
    // NewTemplates sets the config for the template package
    func NewTemplates(a *config.AppConfig) {
        app = a
    }
    
    func AddDefaultData(td *models.TemplateData) *models.TemplateData {
        return td
    }
    
    func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
    
        var tc map[string]*template.Template
        if app.UseCache {
            tc = app.TemplateCache
        } else {
            tc, _ = CreateTemplateCache()
        }
    
        // get request template from cache
        t, ok := tc[tmpl]
        if !ok {
            log.Fatalln("template not found")
        }
    
        buf := new(bytes.Buffer)
    
        td = AddDefaultData(td)
    
        if err := t.Execute(buf, td); err != nil {
            log.Fatalln(err)
        }
    
        // render the template
        _, err := buf.WriteTo(w)
        if err != nil {
            log.Fatalln(err)
        }
    
    }
    
    func CreateTemplateCache() (map[string]*template.Template, error) {
        myCache := map[string]*template.Template{}
        pages, err := filepath.Glob("../../templates/*.page.tmpl")
        if err != nil {
            return myCache, err
        }
    
        for _, page := range pages {
            name := filepath.Base(page)
            ts, err := template.New(name).ParseFiles(page)
            if err != nil {
                return myCache, err
            }
    
            matches, err := filepath.Glob("../../templates/*.layout.tmpl")
            if err != nil {
                return myCache, err
            }
    
            if len(matches) > 0 {
                ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
                if err != nil {
                    return myCache, err
                }
            }
    
            myCache[name] = ts
        }
    
        return myCache, nil
    }

------------------------------------------------------------------------------------------------
    const portNumber = ":8080"
    
    func main() {
        var app config.AppConfig
        tc, err := render.CreateTemplateCache()
        if err != nil {
            log.Fatalln("cannot create template cache")
        }
    
        app.TemplateCache = tc
        app.UseCache = false
    
        repo := handlers.NewRepository(&app)
        handlers.NewHandlers(repo)
    
        render.NewTemplates(&app)
    
        fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
    
        server := &http.Server{
            Addr:    portNumber,
            Handler: route(&app),
        }
    
        err = server.ListenAndServe()
        if err != nil {
            log.Fatalln(err)
        }
    }
