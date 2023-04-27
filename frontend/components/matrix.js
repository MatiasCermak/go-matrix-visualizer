export default function Matrix({ matrixData }) {
    function renderMatrix() {
        return (
            <table className="table table-bordered">
                <tbody>
                    {matrixData.map((row, i) => (
                        <tr key={i}>
                            {row.map((num, j) => (
                                <td key={j}>{num}</td>
                            ))}
                        </tr>
                    ))}
                </tbody>
            </table>
        );
    }

    return matrixData ? renderMatrix() : <div>Matrix will be shown here when submitted.</div>;
}
