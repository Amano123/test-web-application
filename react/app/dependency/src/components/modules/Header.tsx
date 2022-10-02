import React from "react";

export default function Header() {
    return (
        <header
            style={{
                position: "sticky",
                top: 0,
                width: "100%",
                height: 80,
            }}
        >
            <div
                style={{
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "space-between",
                    width: "90%",
                    maxWidth: 900,
                    height: "100%",
                    margin: "0 auto",
                }}
            >
                <h1>
                    <a href="./" style={{ width: "auto", height: "1em" }}>
                        <div>
                            sample page
                        </div>
                    </a>
                </h1>
                <nav
                    style={{
                        display: "flex",
                        alignItems: "center",
                        gap: 20,
                    }}
                >
                    <a
                        href="https://deno.land/x/aleph"
                        style={{ fontSize: 20, color: "#454545" }}
                    >
                        <div>
                            sample page
                        </div>
                    </a>
                    <a
                        href="https://github.com/alephjs/aleph.js"
                        style={{ fontSize: 20, color: "#454545" }}
                    >
                        <div>
                            sample page
                        </div>
                    </a>
                </nav>
            </div>
        </header >
    );
}
