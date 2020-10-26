import Head from 'next/head'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Waypoint Next.js Example</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      
      <header>
        <a href="https://waypointproject.io" className={styles.logo}>
          <img src="/logo.svg" alt="Logo" />
        </a>
      </header>
      <section className={styles.content}>
        <div className={styles.language}>
          <img src="/language.svg" alt="Next.js Icon" />
        </div>
        <h1>This Next.js app was deployed with Waypoint.</h1>
        <p>
          Try making a change to this text locally and run <code>waypoint up</code> again to see it.
        </p>
        <p>
          Read the <a href="https://waypointproject.io/docs">documentation</a> for more about Waypoint.
        </p>
      </section>
      <footer>
        <a href="https://hashicorp.com" className={styles.hashi}>
          <img src="/hashi.svg" alt="HashiCorp" />
        </a>
      </footer>
    </div>
  )
}
