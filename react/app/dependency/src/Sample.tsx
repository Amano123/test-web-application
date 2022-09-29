import React, { useEffect, useState } from "react";

const Sample = () => {
    const [advice, setAdvice] = useState("");
    useEffect(() => {
        // URLを指定
        const api_url = "http://localhost:3000";
        // ここでURLにアクセス
        const fetchData = async () => {
            try {
                const response = await fetch(api_url);
                const json = await response.json();
                // console.log(json);
                setAdvice(json.name);
            } catch (error) {
                console.log("error", error);
                setAdvice("errer");
            }
        };
        fetchData()
    }, [])
    return (
        <div>
            {advice}
        </div>
    );
}

export default Sample;

