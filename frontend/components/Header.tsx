import Link from 'next/link'

const Header = () => {
  return (
    <header className="bg-blue-600 text-white p-4">
      <nav className="container mx-auto flex justify-between">
        <div className="font-bold text-xl">
          <Link href="/">Chacha</Link>
        </div>
        <div>
          <Link href="/login" className="mr-4">Login</Link>
          <Link href="/register">Register</Link>
        </div>
      </nav>
    </header>
  )
}

export default Header
