import './App.css';
import Header from './components/header/Header';
import Footer from './components/footer/Footer';
import Main from './components/main/Main';

function App() {
  return (
    <div className="App">
      <Header title="Users Listing" subtitle="subtitle2"/>
      <Main/>
      <Footer note="Footer note"/>
    </div>
  );
}

export default App;
