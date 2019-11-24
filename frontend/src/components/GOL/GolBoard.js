import React, { useRef, useEffect } from "react";
import PropTypes from "prop-types";
import * as d3 from "d3";
export default function GolBoard(props) {
  //const refElement = useRef(null);
  const boardContainer = useRef(null);
  console.log(props);
  //
  //
  //
  const handleClick = (x, y) => {
    console.log("clicked ->");

    console.log("clicked =>", x, y, props.myColor);
    const move = {
      msgType: "GOLGAME",
      x,
      y,
      playerColor: props.myColor,
      generation: 99,
      playerID: "007",
      //moveType: "instill",
      //text: "",
      payload: "INSTILL"
    };
    props.send(JSON.stringify(move));
    // send click value to server here
  };

  ///////////////////////////////////////////
  //
  //
  ////////////////////////////////////////////
  useEffect(() => {
    //console.log(boardContainer);
    //console.log("----------------->", props.gridData);
    if (props.boardData && boardContainer.current) {
      //console.log("==================>", props.gridData);
      //console.log("generating board");
      const grid = d3
        .select(boardContainer.current)
        //var grid = d3
        // .select("#grid")
        .append("svg")
        .attr("width", "510px")
        .attr("height", "510px");

      var row = grid
        .selectAll(".row")
        .data(props.boardData)
        .enter()
        .append("g")
        .attr("class", "row");

      var column = row
        .selectAll(".square")
        .data(function(d) {
          return d;
        })
        .enter()
        .append("rect")
        .attr("class", "square")
        .attr("x", function(d) {
          return d.x * 50 + 1;
        })
        .attr("y", function(d) {
          return d.y * 50 + 1; //d.y;
        })
        .attr("width", function(d) {
          return 50; //d.width;
        })
        .attr("height", function(d) {
          return 50; //d.height;
        })
        .style("fill", "#fff")
        .style("stroke", "#222")

        .on("click", function(d) {
          d.status = true;

          handleClick(d.x, d.y);
          console.log("clicked");
          //          handleClick.bind(this, "COORDS");
          //this.handleClick("COORDS");
          if (d.status === true) {
            d3.select(this).style("fill", "#0000ff");
          }
        });

      column.style("fill", function(d) {
        if (d.colorHex) {
          return d.colorHex;
        } else return "#fff";
      });

      // grid.style("fill", function(d) {
      //   if (d.colorHex) {
      //     return d.colorHex;
      //   } else return "#fff";
      // });

      grid.exit().remove();
    }
    //}, [props.gridData, boardContainer.current]);
  }, [props.boardData, boardContainer.current]);

  ////////////////////////
  ////////////////////////////////

  return (
    <div>
      GOL BOARD
      <hr />
      <hr />
      <svg
        style={{ height: "100vh", width: "100vw" }}
        className="board-component"
        ref={boardContainer}
      />
    </div>
  );
}

GolBoard.propTypes = {
  boardData: PropTypes.object.isRequired
};
