import React from "react";
import { Container } from "react-bootstrap";

import "pages/about/AboutPage.css";

function AboutPage() {
  return (
    <Container className="about-page">
      <h3>About</h3>

      <div className="trailer-video">
        <div className="trailer-video-embed">
          <iframe
            src="https://www.youtube.com/embed/hNc6wDyovMY"
            title="Shitcamp trailer"
            frameBorder="0"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
            allowFullScreen={true}
          ></iframe>
        </div>
      </div>
    </Container>
  );
}

export default AboutPage;
