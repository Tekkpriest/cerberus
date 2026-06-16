import './App.css'
import './components/navbarmodule/navbar'

function App() {
  return (
    <>
      <header className='app-header'>
        <div className="brand-logo">
          <img src="/Generiertes Bild 2 (3).png" alt="Cerberus Logo" />
        </div>

        <nav className="navbar">
          <a>Dashboard</a>
          <a>Logs</a>
          <a>Settings</a>
        </nav>
      </header>

      <main className="main-hub">
        <div className="watermark">
        <h1>Cerberus</h1>
        <p>control your server!</p>
        </div>

        <div className="main-container">
        </div>
      </main>
    </>
  )
}

export default App