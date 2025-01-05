function showPrinciples() {
    const principles = [
        "Satisfy the customer through early and continuous delivery of valuable software.",
        "Welcome changing requirements, even late in development.",
        "Deliver working software frequently, with a preference for shorter timescales.",
        "Business people and developers must work together daily.",
        "Build projects around motivated individuals. Give them the environment and support they need.",
        "Face-to-face conversation is the most efficient and effective method of conveying information.",
        "Working software is the primary measure of progress.",
        "Agile processes promote sustainable development. Teams should maintain a constant pace indefinitely.",
        "Continuous attention to technical excellence and good design enhances agility.",
        "Simplicity—the art of maximizing the amount of work not done—is essential.",
        "The best architectures, requirements, and designs emerge from self-organizing teams.",
        "At regular intervals, the team reflects on how to become more effective and tunes its behavior."
    ];

    let content = "<h3>Agile Principles</h3><ul>";
    principles.forEach((principle) => {
        content += `<li>${principle}</li>`;
    });
    content += "</ul>";

    const section = document.querySelector("main");
    section.innerHTML += content;
}
