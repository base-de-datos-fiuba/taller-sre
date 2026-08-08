import { FormEvent, useCallback, useEffect, useState } from 'react'

type Product = {
  id: number
  name: string
  description: string
  price: string
  stock: number
}

type DatabaseState = 'checking' | 'online' | 'error'

async function responseError(response: Response): Promise<string> {
  const body = await response.json().catch(() => null) as { error?: string } | null
  return body?.error ?? `Request failed with status ${response.status}`
}

export default function App() {
  const [products, setProducts] = useState<Product[]>([])
  const [search, setSearch] = useState('')
  const [database, setDatabase] = useState<DatabaseState>('checking')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const checkHealth = useCallback(async () => {
    try {
      const response = await fetch('/api/health')
      if (!response.ok) throw new Error(await responseError(response))
      setDatabase('online')
    } catch (requestError) {
      setDatabase('error')
      setError(requestError instanceof Error ? requestError.message : 'Database unavailable')
    }
  }, [])

  const loadProducts = useCallback(async (term = '') => {
    setLoading(true)
    setError('')
    try {
      const query = term ? `?search=${encodeURIComponent(term)}` : ''
      const response = await fetch(`/api/products${query}`)
      if (!response.ok) throw new Error(await responseError(response))
      setProducts(await response.json() as Product[])
      await checkHealth()
    } catch (requestError) {
      setProducts([])
      setDatabase('error')
      setError(requestError instanceof Error ? requestError.message : 'Could not load products')
    } finally {
      setLoading(false)
    }
  }, [checkHealth])

  useEffect(() => {
    void loadProducts()
    const healthTimer = window.setInterval(() => void checkHealth(), 3000)
    return () => window.clearInterval(healthTimer)
  }, [checkHealth, loadProducts])

  function submitSearch(event: FormEvent) {
    event.preventDefault()
    void loadProducts(search)
  }

  return (
    <main>
      <header>
        <div>
          <p className="eyebrow">DATABASE SYSTEMS LAB</p>
          <h1>ByteMarket</h1>
          <p className="subtitle">Simple technology for curious people.</p>
        </div>
        <div className={`status status-${database}`} role="status">
          <span className="status-dot" />
          Database: {database.toUpperCase()}
        </div>
      </header>

      <form className="search" onSubmit={submitSearch}>
        <label htmlFor="product-search">Search the catalog</label>
        <div className="search-row">
          <input
            id="product-search"
            value={search}
            onChange={(event) => setSearch(event.target.value)}
            placeholder="Search products..."
            autoComplete="off"
          />
          <button type="submit" disabled={loading}>{loading ? 'Searching…' : 'Search'}</button>
        </div>
      </form>

      {error && <div className="error" role="alert"><strong>Backend error</strong><span>{error}</span></div>}

      <section aria-live="polite">
        <div className="results-heading">
          <h2>Products</h2>
          <span>{products.length} item{products.length === 1 ? '' : 's'}</span>
        </div>
        <div className="product-grid">
          {products.map((product) => (
            <article className="product-card" key={product.id}>
              <div className="product-icon" aria-hidden="true">⌁</div>
              <h3>{product.name}</h3>
              <p>{product.description}</p>
              <div className="product-meta">
                <strong>${Number(product.price).toFixed(2)}</strong>
                <span>{product.stock} in stock</span>
              </div>
            </article>
          ))}
        </div>
        {!loading && !error && products.length === 0 && <p className="empty">No products found.</p>}
      </section>

      <footer>Disposable classroom environment · never deploy this application</footer>
    </main>
  )
}
