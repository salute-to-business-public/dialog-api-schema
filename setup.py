import os
import setuptools


def readme():
    with open('README.md') as f:
        return f.read()


def version():
    return os.getenv('CIRCLE_TAG')


setuptools.setup(
    name="dialog-api",
    version=version(),
    author="Dialog LLC",
    author_email="services@dlg.im",
    description="Dialog API for Python",
    long_description=readme(),
    long_description_content_type="text/markdown",
    url="https://github.com/dialogs/api-schema",
    packages=setuptools.find_packages(),
    license='Apache License 2.0',
    keywords='dialog messenger',
    python_requires='>=3',
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: Apache Software License",
        "Operating System :: OS Independent",
    ],
)
