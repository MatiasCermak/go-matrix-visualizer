import * as Yup from "yup";

import { Field, Form, Formik } from "formik";

import Matrix from "@/components/matrix";
import calculateMatrix from "@/api/matrix";
import { useState } from "react";

export default function FibonacciSpiral() {
    const [matrix, setMatrix] = useState(null);
    return (
        <div className="container-md vh-100 d-flex flex-column justify-content-center">
            <div className="row">
                <div className="col-3 pb-1 d-flex align-content-center justify-content-start">
                    <h1>Fibonacci Spiral</h1>
                </div>
            </div>
            <div className="col-12">
                <Formik
                    initialValues={{ rows: 1, columns: 1 }}
                    validationSchema={Yup.object().shape({
                        rows: Yup.number()
                            .required("Please enter a number.")
                            .min(1, "The number can't be less than 1."),
                        columns: Yup.number()
                            .required("Please enter a number.")
                            .min(1, "The number can't be less than 1."),
                    })}
                    onSubmit={({ rows, columns }) => {
                        calculateMatrix(rows, columns).then((data) => {
                            if (data.message) {
                                alert(data.message);
                            } else {
                                setMatrix(data.rows);
                            }
                        });
                    }}
                >
                    {({ errors }) => (
                        <Form>
                            <div>Matrix Properties</div>
                            <div class="input-group mb-3">
                                <label for="rows" class="input-group-text">
                                    Rows
                                </label>
                                <Field className="form-control" type="number" name="rows" min="1" />
                                <label for="columns" class="input-group-text">
                                    Columns
                                </label>
                                <Field className="form-control" type="number" name="columns" min="1" />
                                <button class="btn btn-secondary" type="submit">
                                    Calculate
                                </button>
                            </div>
                            {errors.rows ? <div class="alert alert-danger">{errors.rows}</div> : null}
                            {errors.columns && errors.columns !== errors.rows ? (
                                <div class="alert alert-danger">{errors.columns}</div>
                            ) : null}
                        </Form>
                    )}
                </Formik>
            </div>
            <div className="col-12 d-flex align-content-center justify-content-center">
                <Matrix matrixData={matrix} />
            </div>
        </div>
    );
}
