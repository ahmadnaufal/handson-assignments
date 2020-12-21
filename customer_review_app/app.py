from flask import Flask, render_template, request, url_for, redirect
from flask_bootstrap import Bootstrap
from textblob import TextBlob, Word
import random
import time

app = Flask(__name__)
Bootstrap(app)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/analyze', methods=['POST'])
def analyze():
    start_time = time.time()

    if request.method == 'POST':
        raw_text = request.form['raw_text']
        blob = TextBlob(raw_text)
        # get sentiment from blob
        blob_sentiment, blob_subjectivity = blob.sentiment.polarity, blob.sentiment.subjectivity
        n_tokens = len(list(blob.words))

        # fetch nouns
        nouns = []
        summary = []
        for word, tag in blob.tags:
            if tag == 'NN':
                nouns.append(word.lemmatize())
                rand_words = random.sample(nouns, len(nouns))
                for nn in rand_words:
                    summary.append(Word(nn).pluralize())

        duration = time.time() - start_time
        return render_template('index.html', received_text=raw_text, n_tokens=n_tokens, n_words=len(nouns), sentiment_score=blob_sentiment, subjectivity_score=blob_subjectivity, summary=summary, duration=duration)

    return redirect(url_for('index'))
