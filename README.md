<h1 align="center">Project Nyx</h1>
<h2 align="center"><s>Class level: Omega</s></h2>
<p>Nyx is a blazing fast Go-powered chess engine fused with a PyTorch-based ML Layer
for smart move predictions and advanced evaluation</p>
<h3>What is Project Nyx Exactly?</h3>
<p>Project Nyx is a hybrid chess engine built for speed, simplicity and efficiency:
    <ul>
        <li>Go handles the engine core. Move generation, board representation,
        evaluation and legality checking.</li>
        <li>Python, specifically PyTorch powers the ML model. Its trained to suggest
        moves based on game state, learning from classical and modern playstyles</li>
    </ul>
This engine is designed to be lightweight, extensible, efficient and smart enough
to give you a real challenge, and get smarter by learning from your games.
</p>
<h3>Features:</h3>
<ul>
    <li>Engine written in pure Go for its performance and lightweight concurrency model</li>
    <li>Integrated PyTorch model for position evaluation and optimal move suggestion</li>
    <li>Self-play + training loop support</li>
    <li>Load/save game state via FEN/PGN</li>
</ul>
<h3> Prerequisites: </h3>
<ul>
    <li>Python 3.10+</li>
    <li>Go 1.20+</li>
    <li>PyTorch (with CUDA support if available)</li>
</ul>
<h3>How the AI works</h3>
<h4>The engine sends board states to the ML model via shared memory. The PyTorch
model returns evaluation scores and candidate moves, which the Go engine uses in its
search algorithm.<br> You can train the model on:
    <ul>
        <li>Historical PGN datasets (Grandmaster games, Lichess, etc.)</li>
        <li>Self-play games using reinforcement learning</li>
    </ul>
</h4>
