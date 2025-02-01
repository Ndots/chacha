import type { NextPage } from 'next'
import Header from '../components/Header'
import Footer from '../components/Footer'

const Home: NextPage = () => {
  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-4xl font-bold mb-4">Welcome to Chacha</h1>
        <p className="text-lg">The one-stop platform for business registration.</p>
      </main>
      <Footer />
    </div>
  )
}

export default Home
