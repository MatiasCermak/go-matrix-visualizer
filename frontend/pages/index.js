import Link from "next/link";

export default function Home() {
    return (
        <>
            <header className="banner d-flex align-content-center justify-content-center">
                <h1>¡Bienvenido!</h1>
            </header>
            <div className="d-flex align-content-center justify-content-center">
                <Link href="/matrix/fibonacciSpiral">
                    <div className="btn btn-primary">Acceder a App</div>
                </Link>
            </div>
        </>
    );
}
